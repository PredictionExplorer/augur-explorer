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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"}],\"name\":\"EthDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"wasTokenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 tokenId_) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTCaller) WasTokenUsed(opts *bind.CallOpts, tokenId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _IStakingWalletCST.contract.Call(opts, &out, "wasTokenUsed", tokenId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 tokenId_) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTSession) WasTokenUsed(tokenId_ *big.Int) (bool, error) {
	return _IStakingWalletCST.Contract.WasTokenUsed(&_IStakingWalletCST.CallOpts, tokenId_)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 tokenId_) view returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTCallerSession) WasTokenUsed(tokenId_ *big.Int) (bool, error) {
	return _IStakingWalletCST.Contract.WasTokenUsed(&_IStakingWalletCST.CallOpts, tokenId_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) DepositIfPossible(opts *bind.TransactOpts, roundNum_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "depositIfPossible", roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.DepositIfPossible(&_IStakingWalletCST.TransactOpts, roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.DepositIfPossible(&_IStakingWalletCST.TransactOpts, roundNum_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 tokenId_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) Stake(opts *bind.TransactOpts, tokenId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "stake", tokenId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 tokenId_) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) Stake(tokenId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Stake(&_IStakingWalletCST.TransactOpts, tokenId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 tokenId_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) Stake(tokenId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Stake(&_IStakingWalletCST.TransactOpts, tokenId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] tokenIds_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) StakeMany(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "stakeMany", tokenIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] tokenIds_) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) StakeMany(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.StakeMany(&_IStakingWalletCST.TransactOpts, tokenIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] tokenIds_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) StakeMany(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.StakeMany(&_IStakingWalletCST.TransactOpts, tokenIds_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTTransactor) TryPerformMaintenance(opts *bind.TransactOpts, resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "tryPerformMaintenance", resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.TryPerformMaintenance(&_IStakingWalletCST.TransactOpts, resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.TryPerformMaintenance(&_IStakingWalletCST.TransactOpts, resetState_, charityAddress_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Unstake(&_IStakingWalletCST.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.Unstake(&_IStakingWalletCST.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletCST *IStakingWalletCSTSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeMany(&_IStakingWalletCST.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletCST *IStakingWalletCSTTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletCST.Contract.UnstakeMany(&_IStakingWalletCST.TransactOpts, stakeActionIds_)
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
	RoundNum      *big.Int
	ActionCounter *big.Int
	DepositIndex  *big.Int
	DepositId     *big.Int
	DepositAmount *big.Int
	NumStakedNFTs *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthDepositEvent is a free log retrieval operation binding the contract event 0x725e2879ec3f698d2b2186dc65d2b3c8f4768f1fd80c31be29fe310a83500c98.
//
// Solidity: event EthDepositEvent(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNFTs)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) FilterEthDepositEvent(opts *bind.FilterOpts, roundNum []*big.Int) (*IStakingWalletCSTEthDepositEventIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.FilterLogs(opts, "EthDepositEvent", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTEthDepositEventIterator{contract: _IStakingWalletCST.contract, event: "EthDepositEvent", logs: logs, sub: sub}, nil
}

// WatchEthDepositEvent is a free log subscription operation binding the contract event 0x725e2879ec3f698d2b2186dc65d2b3c8f4768f1fd80c31be29fe310a83500c98.
//
// Solidity: event EthDepositEvent(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNFTs)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) WatchEthDepositEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletCSTEthDepositEvent, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.WatchLogs(opts, "EthDepositEvent", roundNumRule)
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

// ParseEthDepositEvent is a log parse operation binding the contract event 0x725e2879ec3f698d2b2186dc65d2b3c8f4768f1fd80c31be29fe310a83500c98.
//
// Solidity: event EthDepositEvent(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNFTs)
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
	StakeActionId *big.Int
	TokenId       *big.Int
	StakerAddress common.Address
	NumStakedNFTs *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0xcd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef8.
//
// Solidity: event StakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (*IStakingWalletCSTStakeActionEventIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.FilterLogs(opts, "StakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTStakeActionEventIterator{contract: _IStakingWalletCST.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0xcd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef8.
//
// Solidity: event StakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletCSTStakeActionEvent, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.WatchLogs(opts, "StakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
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

// ParseStakeActionEvent is a log parse operation binding the contract event 0xcd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef8.
//
// Solidity: event StakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs)
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
	StakeActionId *big.Int
	TokenId       *big.Int
	StakerAddress common.Address
	NumStakedNFTs *big.Int
	RewardAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c96.
//
// Solidity: event UnstakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs, uint256 rewardAmount)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (*IStakingWalletCSTUnstakeActionEventIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.FilterLogs(opts, "UnstakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletCSTUnstakeActionEventIterator{contract: _IStakingWalletCST.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c96.
//
// Solidity: event UnstakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs, uint256 rewardAmount)
func (_IStakingWalletCST *IStakingWalletCSTFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *IStakingWalletCSTUnstakeActionEvent, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _IStakingWalletCST.contract.WatchLogs(opts, "UnstakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
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

// ParseUnstakeActionEvent is a log parse operation binding the contract event 0x678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c96.
//
// Solidity: event UnstakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs, uint256 rewardAmount)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"game_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"}],\"name\":\"AccessError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"DepositFromUnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"FundTransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"InvalidOperationInCurrentState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"NoTokensStaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OneTimeStaking\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"}],\"name\":\"TokenAlreadyUnstaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"}],\"name\":\"EthDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"FundTransferFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"charityAddress\",\"type\":\"address\"}],\"name\":\"FundsTransferredToCharityEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"}],\"name\":\"StakeActionEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNFTs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"UnstakeActionEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ETHDepositIndex\",\"type\":\"uint256\"}],\"name\":\"ETHDeposits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"depositId\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"rewardAmountPerStakedNFT\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundNum_\",\"type\":\"uint256\"}],\"name\":\"depositIfPossible\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"game\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractCosmicSignature\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numETHDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numTokensStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"resetState_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"charityAddress_\",\"type\":\"address\"}],\"name\":\"tryPerformMaintenance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"wasTokenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052346100305761001a610014610133565b906103f7565b610022610035565b61213d61058b823961213d90f35b61003b565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906100679061003f565b810190811060018060401b0382111761007f57604052565b610049565b90610097610090610035565b928361005d565b565b5f80fd5b60018060a01b031690565b6100b19061009d565b90565b6100bd906100a8565b90565b6100c9816100b4565b036100d057565b5f80fd5b905051906100e1826100c0565b565b6100ec816100a8565b036100f357565b5f80fd5b90505190610104826100e3565b565b919060408382031261012e578061012261012b925f86016100d4565b936020016100f7565b90565b610099565b6101516126c88038038061014681610084565b928339810190610106565b9091565b5f1b90565b906101665f1991610155565b9181191691161790565b90565b90565b90565b61018d61018861019292610170565b610176565b610173565b90565b90565b906101ad6101a86101b492610179565b610195565b825461015a565b9055565b6101cc6101c76101d19261009d565b610176565b61009d565b90565b6101dd906101b8565b90565b6101e9906101d4565b90565b90565b6102036101fe610208926101ec565b610176565b61009d565b90565b610214906101ef565b90565b60209181520190565b60207f66742e0000000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f7220746865206e5f8201520152565b61027a6023604092610217565b61028381610220565b0190565b61029c9060208101905f81830391015261026d565b90565b156102a657565b6102ae610035565b63eac0d38960e01b8152806102c560048201610287565b0390fd5b60207f616d652e00000000000000000000000000000000000000000000000000000000917f5a65726f2d616464726573732077617320676976656e20666f722074686520675f8201520152565b6103236024604092610217565b61032c816102c9565b0190565b6103459060208101905f818303910152610316565b90565b1561034f57565b610357610035565b63eac0d38960e01b81528061036e60048201610330565b0390fd5b9061038360018060a01b0391610155565b9181191691161790565b610396906101b8565b90565b6103a29061038d565b90565b90565b906103bd6103b86103c492610399565b6103a5565b8254610372565b9055565b6103d1906101d4565b90565b90565b906103ec6103e76103f3926103c8565b6103d4565b8254610372565b9055565b9061046d6104749261040833610476565b61041460026005610198565b610441610420826101e0565b61043a61043461042f5f61020b565b6100a8565b916100a8565b141561029f565b6104668361045f6104596104545f61020b565b6100a8565b916100a8565b1415610348565b60016103a8565b60026103d7565b565b61047f906104a3565b565b61048a906100a8565b9052565b91906104a1905f60208501940190610481565b565b806104be6104b86104b35f61020b565b6100a8565b916100a8565b146104ce576104cc9061052b565b565b6104f16104da5f61020b565b5f918291631e4fbdf760e01b83526004830161048e565b0390fd5b5f1c90565b60018060a01b031690565b610511610516916104f5565b6104fa565b90565b6105239054610505565b90565b5f0190565b6105345f610519565b61053e825f6103d7565b9061057261056c7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936103c8565b916103c8565b9161057b610035565b8061058581610526565b0390a356fe60806040526004361015610013575b6109d1565b61001d5f3561011c565b80630d50c1891461011757806315b4e68f146101125780632a3247aa1461010d5780632e17de7814610108578063451f1adf1461010357806347ccca02146100fe57806355279fdb146100f95780635fda0acc146100f4578063715018a6146100ef5780638da5cb5b146100ea578063a2b136fb146100e5578063a694fc3a146100e0578063c3fe3e28146100db578063d8ee5573146100d6578063f2fde38b146100d15763fe939afc0361000e5761099e565b61096b565b610917565b61086f565b610807565b6107d1565b6106f0565b61068f565b61065a565b610625565b610596565b6104ac565b610365565b610330565b6102e0565b61028f565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061016090610138565b810190811067ffffffffffffffff82111761017a57604052565b610142565b9061019261018b610122565b9283610156565b565b67ffffffffffffffff81116101ac5760208091020190565b610142565b5f80fd5b90565b6101c1816101b5565b036101c857565b5f80fd5b905035906101d9826101b8565b565b909291926101f06101eb82610194565b61017f565b938185526020808601920283019281841161022d57915b8383106102145750505050565b6020809161022284866101cc565b815201920191610207565b6101b1565b9080601f830112156102505781602061024d933591016101db565b90565b610134565b90602082820312610285575f82013567ffffffffffffffff81116102805761027d9201610232565b90565b610130565b61012c565b5f0190565b346102bd576102a76102a2366004610255565b610a81565b6102af610122565b806102b98161028a565b0390f35b610128565b906020828203126102db576102d8915f016101cc565b90565b61012c565b6102f36102ee3660046102c2565b610f8a565b6102fb610122565b806103058161028a565b0390f35b151590565b61031790610309565b9052565b919061032e905f6020850194019061030e565b565b346103605761035c61034b6103463660046102c2565b6111b4565b610353610122565b9182918261031b565b0390f35b610128565b346103935761037d6103783660046102c2565b6111d3565b610385610122565b8061038f8161028a565b0390f35b610128565b90565b6103af6103aa6103b4926101b5565b610398565b6101b5565b90565b906103c19061039b565b5f5260205260405f2090565b5f1c90565b67ffffffffffffffff1690565b6103eb6103f0916103cd565b6103d2565b90565b6103fd90546103df565b90565b60401c90565b60018060c01b031690565b61041d61042291610400565b610406565b90565b61042f9054610411565b90565b61043d9060076103b7565b906104545f61044d8185016103f3565b9301610425565b90565b67ffffffffffffffff1690565b61046d90610457565b9052565b60018060c01b031690565b61048590610471565b9052565b9160206104aa9294936104a360408201965f830190610464565b019061047c565b565b346104dd576104c46104bf3660046102c2565b610432565b906104d96104d0610122565b92839283610489565b0390f35b610128565b5f9103126104ec57565b61012c565b1c90565b60018060a01b031690565b61051090600861051593026104f1565b6104f5565b90565b906105239154610500565b90565b61053260015f90610518565b90565b60018060a01b031690565b61055461054f61055992610535565b610398565b610535565b90565b61056590610540565b90565b6105719061055c565b90565b61057d90610568565b9052565b9190610594905f60208501940190610574565b565b346105c6576105a63660046104e2565b6105c26105b1610526565b6105b9610122565b91829182610581565b0390f35b610128565b90565b6105de9060086105e393026104f1565b6105cb565b90565b906105f191546105ce565b90565b61060060085f906105e6565b90565b61060c906101b5565b9052565b9190610623905f60208501940190610603565b565b34610655576106353660046104e2565b6106516106406105f4565b610648610122565b91829182610610565b0390f35b610128565b3461068a5761066a3660046104e2565b6106866106756111ea565b61067d610122565b91829182610610565b0390f35b610128565b346106bd5761069f3660046104e2565b6106a761124d565b6106af610122565b806106b98161028a565b0390f35b610128565b6106cb90610535565b90565b6106d7906106c2565b9052565b91906106ee905f602085019401906106ce565b565b34610720576107003660046104e2565b61071c61070b61125b565b610713610122565b918291826106db565b0390f35b610128565b9061072f9061039b565b5f5260205260405f2090565b61074761074c916103cd565b6105cb565b90565b610759905461073b565b90565b60018060a01b031690565b610773610778916103cd565b61075c565b90565b6107859054610767565b90565b610793906003610725565b906107ab60016107a45f850161074f565b930161077b565b90565b9160206107cf9294936107c860408201965f830190610603565b01906106ce565b565b34610802576107e96107e43660046102c2565b610788565b906107fe6107f5610122565b928392836107ae565b0390f35b610128565b346108355761081f61081a3660046102c2565b611515565b610827610122565b806108318161028a565b0390f35b610128565b61084a90600861084f93026104f1565b61075c565b90565b9061085d915461083a565b90565b61086c60025f90610852565b90565b3461089f5761087f3660046104e2565b61089b61088a610860565b610892610122565b918291826106db565b0390f35b610128565b6108ad81610309565b036108b457565b5f80fd5b905035906108c5826108a4565b565b6108d0816106c2565b036108d757565b5f80fd5b905035906108e8826108c7565b565b9190604083820312610912578061090661090f925f86016108b8565b936020016108db565b90565b61012c565b346109485761094461093361092d3660046108ea565b9061198a565b61093b610122565b9182918261031b565b0390f35b610128565b9060208282031261096657610963915f016108db565b90565b61012c565b346109995761098361097e36600461094d565b611a04565b61098b610122565b806109958161028a565b0390f35b610128565b346109cc576109b66109b1366004610255565b611a0f565b6109be610122565b806109c88161028a565b0390f35b610128565b5f80fd5b90565b6109ec6109e76109f1926109d5565b610398565b6101b5565b90565b6001610a0091016101b5565b90565b5190565b634e487b7160e01b5f52603260045260245ffd5b90610a2582610a03565b811015610a36576020809102010190565b610a07565b610a4590516101b5565b90565b634e487b7160e01b5f52601160045260245ffd5b610a6b610a71919392936101b5565b926101b5565b8201809211610a7c57565b610a48565b90610a8b5f6109d8565b610a945f6109d8565b905b81610ab1610aab610aa687610a03565b6101b5565b916101b5565b1015610aeb57610adf610ae591610ad9610ad4610acf888790610a1b565b610a3b565b611c80565b90610a5c565b916109f4565b90610a96565b9050610af8919250611f59565b565b60209181520190565b60207f7065726d697474656420746f206d616b652061206465706f7369742e00000000917f4f6e6c792074686520436f736d696347616d6520636f6e7472616374206973205f8201520152565b610b5d603c604092610afa565b610b6681610b03565b0190565b9190610b8d906020610b85604086018681035f880152610b50565b9401906106ce565b565b15610b975750565b610bb990610ba3610122565b918291637ed5977760e11b835260048301610b6a565b0390fd5b5f7f546865726520617265206e6f20435354204e465473207374616b65642e000000910152565b610bf1601d602092610afa565b610bfa81610bbd565b0190565b610c139060208101905f818303910152610be4565b90565b610c20604061017f565b90565b5f90565b5f90565b610c33610c16565b9060208083610c40610c23565b815201610c4b610c27565b81525050565b610c59610c2b565b90565b90565b610c73610c6e610c7892610c5c565b610398565b6101b5565b90565b5f1b90565b90610c8c5f1991610c7b565b9181191691161790565b90565b90610cae610ca9610cb59261039b565b610c96565b8254610c80565b9055565b90565b610cd0610ccb610cd592610cb9565b610398565b6101b5565b90565b90610ce290610457565b9052565b90610cf090610471565b9052565b90610d2a610d215f610d04610c16565b94610d1b610d138383016103f3565b838801610cd8565b01610425565b60208401610ce6565b565b610d3590610cf4565b90565b634e487b7160e01b5f52601260045260245ffd5b610d58610d5e916101b5565b916101b5565b908115610d69570490565b610d38565b610d82610d7d610d87926101b5565b610398565b610471565b90565b610d949051610471565b90565b610da3610da991610471565b91610471565b019060018060c01b038211610dba57565b610a48565b610dc8906101b5565b5f198114610dd65760010190565b610a48565b610def610dea610df4926101b5565b610398565b610457565b90565b634e487b7160e01b5f525f60045260245ffd5b610e149051610457565b90565b90610e2a67ffffffffffffffff91610c7b565b9181191691161790565b610e48610e43610e4d92610457565b610398565b610457565b90565b90565b90610e68610e63610e6f92610e34565b610e50565b8254610e17565b9055565b60401b90565b90610e8d67ffffffffffffffff1991610e73565b9181191691161790565b610eab610ea6610eb092610471565b610398565b610471565b90565b90565b90610ecb610ec6610ed292610e97565b610eb3565b8254610e79565b9055565b90610f0060205f610f0694610ef8828201610ef2848801610e0a565b90610e53565b019201610d8a565b90610eb6565b565b90610f1291610ed6565b565b610f28610f23610f2d92610457565b610398565b6101b5565b90565b610f3990610f14565b9052565b90959492610f8894610f77610f8192610f6d608096610f6360a088019c5f890190610603565b6020870190610603565b6040850190610f30565b6060830190610603565b0190610603565b565b610fb133610fa9610fa3610f9e600261077b565b6106c2565b916106c2565b143390610b8f565b610fbb600461074f565b9081610fcf610fc95f6109d8565b916101b5565b1461115057610fdc610c51565b610fe6600861074f565b91611004610ff4600961074f565b610ffe6001610c5f565b90610a5c565b92611010846009610c99565b61101a600561074f565b61102d6110276002610cbc565b916101b5565b10155f146110fa576110b75f6110596110f59361105461104d6001610c5f565b6005610c99565b610dbf565b93611065856008610c99565b61107961107188610ddb565b838801610cd8565b61109861108f61108a348b90610d4c565b610d6e565b60208801610ce6565b5b6110ae866110a9600788906103b7565b610f08565b95939401610e0a565b94346110e37f725e2879ec3f698d2b2186dc65d2b3c8f4768f1fd80c31be29fe310a83500c989661039b565b966110ec610122565b95869586610f3d565b0390a2565b9091506110f56110b75f611118611113600786906103b7565b610d2c565b9461114b61112f61112a348b90610d4c565b610d6e565b611145602089019161114083610d8a565b610d97565b90610ce6565b611099565b611158610122565b63bc8b155960e01b81528061116f60048201610bfe565b0390fd5b5f90565b906111819061039b565b5f5260205260405f2090565b60ff1690565b61119f6111a4916103cd565b61118d565b90565b6111b19054611193565b90565b6111cb6111d0916111c3611173565b506006611177565b6111a7565b90565b6111df6111e491611c80565b611f59565b565b5f90565b6111f26111e6565b506111fd600461074f565b90565b611208611f8b565b61121061123a565b565b61122661122161122b926109d5565b610398565b610535565b90565b61123790611212565b90565b61124b6112465f61122e565b611fd9565b565b611255611200565b565b5f90565b611263611257565b5061126d5f61077b565b90565b60407f206f6e6c79206f6e63652e000000000000000000000000000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f2e20416e204e465420697320616c6c6f77656420746f206265207374616b656460208201520152565b6112f0604b606092610afa565b6112f981611270565b0190565b9190611320906020611318604086018681035f8801526112e3565b940190610603565b565b1561132a5750565b61134c90611336610122565b918291632290948760e21b8352600483016112fd565b0390fd5b61135a604061017f565b90565b5f90565b5f90565b61136d611350565b906020808361137a61135d565b815201611385611361565b81525050565b611393611365565b90565b906113a0906101b5565b9052565b906113ae906106c2565b9052565b6113bc90516106c2565b90565b906113d060018060a01b0391610c7b565b9181191691161790565b6113e39061055c565b90565b90565b906113fe6113f9611405926113da565b6113e6565b82546113bf565b9055565b906114346020600161143a9461142c5f82016114265f8801610a3b565b90610c99565b0192016113b2565b906113e9565b565b9061144691611409565b565b9061145460ff91610c7b565b9181191691161790565b61146790610309565b90565b90565b9061148261147d6114899261145e565b61146a565b8254611448565b9055565b61149961149e916103cd565b6104f5565b90565b6114ab905461148d565b90565b6114b79061055c565b90565b5f80fd5b60e01b90565b5f9103126114ce57565b61012c565b6040906114fc61150394969593966114f260608401985f8501906106ce565b60208301906106ce565b0190610603565b565b61150d610122565b3d5f823e3d90fd5b61153c61153561152f61152a60068590611177565b6111a7565b15610309565b8290611322565b61154461138b565b611550825f8301611396565b61155d33602083016113a4565b61158f61157261156d600961074f565b610dbf565b61157d816009610c99565b9161158a60038490610725565b61143c565b6115ac61159c600461074f565b6115a66001610c5f565b90610a5c565b906115b8826004610c99565b6115cc6115c56002610cbc565b6005610c99565b6115e260016115dd60068690611177565b61146d565b6115f46115ef60016114a1565b610568565b6323b872dd33611603306114ae565b928692813b156116d1575f61162b91611636829661161f610122565b988997889687956114be565b8552600485016114d3565b03925af180156116cc576116a0575b509133919261169b61168961168361167d7fcd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef89461039b565b9461039b565b946113da565b94611692610122565b91829182610610565b0390a4565b6116bf905f3d81116116c5575b6116b78183610156565b8101906114c4565b5f611645565b503d6116ad565b611505565b6114ba565b906116e992916116e4611f8b565b61185c565b90565b5f7f546865726520617265207374696c6c20435354204e465473207374616b65642e910152565b61171f60208092610afa565b611728816116ec565b0190565b6117419060208101905f818303910152611713565b90565b1561174b57565b611753610122565b63a29f5c4d60e01b81528061176a6004820161172c565b0390fd5b905090565b61177e5f809261176e565b0190565b61178b90611773565b90565b67ffffffffffffffff81116117ac576117a8602091610138565b0190565b610142565b906117c36117be8361178e565b61017f565b918252565b606090565b3d5f146117e8576117dd3d6117b1565b903d5f602084013e5b565b6117f06117c8565b906117e6565b5f7f5472616e7366657220746f2063686172697479206661696c65642e0000000000910152565b61182a601b602092610afa565b611833816117f6565b0190565b919061185a906020611852604086018681035f88015261181d565b940190610603565b565b5061188261186a600461074f565b61187c6118765f6109d8565b916101b5565b14611744565b611972575b806118a261189c6118975f61122e565b6106c2565b916106c2565b036118ae575b50600190565b6118b7306114ae565b315f8083836118c4610122565b90816118cf81611782565b03925af16118db6117cd565b505f1461192b57906119226119107f80c1082d1fcf8195bbf5a158fbef654d58f69408bd2e339b466bbd7c9fd7f74e926113da565b92611919610122565b91829182610610565b0390a25f6118a8565b9061196b6119597fc4283eec4c85bbabb45475fc62cda5e1b3a382be7ef6b1d14e685b065a9adf74926113da565b92611962610122565b91829182611837565b0390a25f90565b61198561197e5f6109d8565b6008610c99565b611887565b9061199c91611997611173565b6116d6565b90565b6119b0906119ab611f8b565b6119b2565b565b806119cd6119c76119c25f61122e565b6106c2565b916106c2565b146119dd576119db90611fd9565b565b611a006119e95f61122e565b5f918291631e4fbdf760e01b8352600483016106db565b0390fd5b611a0d9061199f565b565b90611a195f6109d8565b5b80611a35611a2f611a2a86610a03565b6101b5565b916101b5565b1015611a6457611a5f90611a5a611a55611a50868490610a1b565b610a3b565b611515565b6109f4565b611a1a565b509050565b90611aa0611a976001611a7a611350565b94611a91611a895f830161074f565b5f8801611396565b0161077b565b602084016113a4565b565b611aab90611a69565b90565b5f7f4e46542068617320616c7265616479206265656e20756e7374616b65642e0000910152565b611ae2601e602092610afa565b611aeb81611aae565b0190565b9190611b12906020611b0a604086018681035f880152611ad5565b940190610603565b565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b611b6e602a604092610afa565b611b7781611b14565b0190565b916040611bac929493611ba5611b9a606083018381035f850152611b61565b966020830190610603565b01906106ce565b565b1b90565b91906008611bcd910291611bc75f1984611bae565b92611bae565b9181191691161790565b9190611bed611be8611bf59361039b565b610c96565b908354611bb2565b9055565b611c0b91611c056111e6565b91611bd7565b565b60015f91611c1d83808301611bf9565b0155565b905f03611c3357611c3190611c0d565b565b610df7565b611c47611c4d919392936101b5565b926101b5565b8203918211611c5857565b610a48565b916020611c7e929493611c7760408201965f830190610603565b0190610603565b565b611c886111e6565b50611c9d611c9860038390610725565b611aa2565b33611cbb611cb5611cb0602085016113b2565b6106c2565b916106c2565b03611e1657611cc982612071565b91611cdf5f611cda60038490610725565b611c21565b611cfc611cec600461074f565b611cf66001610c5f565b90611c38565b90611d08826004610c99565b611d1a611d1560016114a1565b610568565b906323b872dd611d29306114ae565b3393611d365f8801610a3b565b92813b15611e11575f611d5c91611d678296611d50610122565b998a97889687956114be565b8552600485016114d3565b03925af1908115611e0c57611d85925f92611de0575b509301610a3b565b33919284611dc5611dbf611db97f678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c969461039b565b9461039b565b946113da565b94611dda611dd1610122565b92839283611c5d565b0390a490565b611dff90833d8111611e05575b611df78183610156565b8101906114c4565b5f611d7d565b503d611ded565b611505565b6114ba565b6020611e2291016113b2565b611e3c611e36611e315f61122e565b6106c2565b916106c2565b14155f14611e6b573390611e67611e51610122565b9283926345c2e43b60e01b845260048401611b7b565b0390fd5b611e8d90611e77610122565b91829163aed59e4f60e01b835260048301611aef565b0390fd5b60207f6661696c65642e00000000000000000000000000000000000000000000000000917f435354204e4654207374616b696e6720726577617264207472616e73666572205f8201520152565b611eeb6027604092610afa565b611ef481611e91565b0190565b916040611f29929493611f22611f17606083018381035f850152611ede565b966020830190610603565b01906106ce565b565b15611f34575050565b611f55611f3f610122565b92839263310a0fbb60e21b845260048401611ef8565b0390fd5b611f89905f803383611f69610122565b9081611f7481611782565b03925af1611f806117cd565b50903391611f2b565b565b611f9361125b565b611fac611fa6611fa16120fa565b6106c2565b916106c2565b03611fb357565b611fd5611fbe6120fa565b5f91829163118cdaa760e01b8352600483016106db565b0390fd5b611fe25f61077b565b611fec825f6113e9565b9061202061201a7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0936113da565b916113da565b91612029610122565b806120338161028a565b0390a3565b600161204491036101b5565b90565b61205b61205661206092610471565b610398565b6101b5565b90565b9061206e91016101b5565b90565b6120796111e6565b506120835f6109d8565b9061208e600861074f565b915b6120a461209f600785906103b7565b610d2c565b6120af5f8201610e0a565b6120c16120bb856101b5565b91610f14565b106120f3576120ed916120e16120dc60206120e79401610d8a565b612047565b90612063565b92612038565b91612090565b5091505090565b612102611257565b50339056fea26469706673582212206063bbf3f995e239f3fefe982a6f2f49f27b2d481ddb02eb90c1269dec82626e64736f6c634300081a0033",
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
// Solidity: function ETHDeposits(uint256 ETHDepositIndex) view returns(uint64 depositId, uint192 rewardAmountPerStakedNFT)
func (_StakingWalletCST *StakingWalletCSTCaller) ETHDeposits(opts *bind.CallOpts, ETHDepositIndex *big.Int) (struct {
	DepositId                uint64
	RewardAmountPerStakedNFT *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "ETHDeposits", ETHDepositIndex)

	outstruct := new(struct {
		DepositId                uint64
		RewardAmountPerStakedNFT *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardAmountPerStakedNFT = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ETHDepositIndex) view returns(uint64 depositId, uint192 rewardAmountPerStakedNFT)
func (_StakingWalletCST *StakingWalletCSTSession) ETHDeposits(ETHDepositIndex *big.Int) (struct {
	DepositId                uint64
	RewardAmountPerStakedNFT *big.Int
}, error) {
	return _StakingWalletCST.Contract.ETHDeposits(&_StakingWalletCST.CallOpts, ETHDepositIndex)
}

// ETHDeposits is a free data retrieval call binding the contract method 0x451f1adf.
//
// Solidity: function ETHDeposits(uint256 ETHDepositIndex) view returns(uint64 depositId, uint192 rewardAmountPerStakedNFT)
func (_StakingWalletCST *StakingWalletCSTCallerSession) ETHDeposits(ETHDepositIndex *big.Int) (struct {
	DepositId                uint64
	RewardAmountPerStakedNFT *big.Int
}, error) {
	return _StakingWalletCST.Contract.ETHDeposits(&_StakingWalletCST.CallOpts, ETHDepositIndex)
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
// Solidity: function stakeActions(uint256 stakeActionId) view returns(uint256 tokenId, address nftOwnerAddress)
func (_StakingWalletCST *StakingWalletCSTCaller) StakeActions(opts *bind.CallOpts, stakeActionId *big.Int) (struct {
	TokenId         *big.Int
	NftOwnerAddress common.Address
}, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "stakeActions", stakeActionId)

	outstruct := new(struct {
		TokenId         *big.Int
		NftOwnerAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftOwnerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 stakeActionId) view returns(uint256 tokenId, address nftOwnerAddress)
func (_StakingWalletCST *StakingWalletCSTSession) StakeActions(stakeActionId *big.Int) (struct {
	TokenId         *big.Int
	NftOwnerAddress common.Address
}, error) {
	return _StakingWalletCST.Contract.StakeActions(&_StakingWalletCST.CallOpts, stakeActionId)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 stakeActionId) view returns(uint256 tokenId, address nftOwnerAddress)
func (_StakingWalletCST *StakingWalletCSTCallerSession) StakeActions(stakeActionId *big.Int) (struct {
	TokenId         *big.Int
	NftOwnerAddress common.Address
}, error) {
	return _StakingWalletCST.Contract.StakeActions(&_StakingWalletCST.CallOpts, stakeActionId)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 tokenId_) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCaller) WasTokenUsed(opts *bind.CallOpts, tokenId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingWalletCST.contract.Call(opts, &out, "wasTokenUsed", tokenId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 tokenId_) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTSession) WasTokenUsed(tokenId_ *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.WasTokenUsed(&_StakingWalletCST.CallOpts, tokenId_)
}

// WasTokenUsed is a free data retrieval call binding the contract method 0x2a3247aa.
//
// Solidity: function wasTokenUsed(uint256 tokenId_) view returns(bool)
func (_StakingWalletCST *StakingWalletCSTCallerSession) WasTokenUsed(tokenId_ *big.Int) (bool, error) {
	return _StakingWalletCST.Contract.WasTokenUsed(&_StakingWalletCST.CallOpts, tokenId_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) DepositIfPossible(opts *bind.TransactOpts, roundNum_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "depositIfPossible", roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_StakingWalletCST *StakingWalletCSTSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.DepositIfPossible(&_StakingWalletCST.TransactOpts, roundNum_)
}

// DepositIfPossible is a paid mutator transaction binding the contract method 0x15b4e68f.
//
// Solidity: function depositIfPossible(uint256 roundNum_) payable returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) DepositIfPossible(roundNum_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.DepositIfPossible(&_StakingWalletCST.TransactOpts, roundNum_)
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
// Solidity: function stake(uint256 tokenId_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) Stake(opts *bind.TransactOpts, tokenId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "stake", tokenId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 tokenId_) returns()
func (_StakingWalletCST *StakingWalletCSTSession) Stake(tokenId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Stake(&_StakingWalletCST.TransactOpts, tokenId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 tokenId_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) Stake(tokenId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Stake(&_StakingWalletCST.TransactOpts, tokenId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] tokenIds_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) StakeMany(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "stakeMany", tokenIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] tokenIds_) returns()
func (_StakingWalletCST *StakingWalletCSTSession) StakeMany(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.StakeMany(&_StakingWalletCST.TransactOpts, tokenIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] tokenIds_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) StakeMany(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.StakeMany(&_StakingWalletCST.TransactOpts, tokenIds_)
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

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_StakingWalletCST *StakingWalletCSTTransactor) TryPerformMaintenance(opts *bind.TransactOpts, resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "tryPerformMaintenance", resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_StakingWalletCST *StakingWalletCSTSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.TryPerformMaintenance(&_StakingWalletCST.TransactOpts, resetState_, charityAddress_)
}

// TryPerformMaintenance is a paid mutator transaction binding the contract method 0xd8ee5573.
//
// Solidity: function tryPerformMaintenance(bool resetState_, address charityAddress_) returns(bool)
func (_StakingWalletCST *StakingWalletCSTTransactorSession) TryPerformMaintenance(resetState_ bool, charityAddress_ common.Address) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.TryPerformMaintenance(&_StakingWalletCST.TransactOpts, resetState_, charityAddress_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletCST *StakingWalletCSTSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Unstake(&_StakingWalletCST.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.Unstake(&_StakingWalletCST.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletCST *StakingWalletCSTSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeMany(&_StakingWalletCST.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletCST *StakingWalletCSTTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletCST.Contract.UnstakeMany(&_StakingWalletCST.TransactOpts, stakeActionIds_)
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
	RoundNum      *big.Int
	ActionCounter *big.Int
	DepositIndex  *big.Int
	DepositId     *big.Int
	DepositAmount *big.Int
	NumStakedNFTs *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthDepositEvent is a free log retrieval operation binding the contract event 0x725e2879ec3f698d2b2186dc65d2b3c8f4768f1fd80c31be29fe310a83500c98.
//
// Solidity: event EthDepositEvent(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNFTs)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterEthDepositEvent(opts *bind.FilterOpts, roundNum []*big.Int) (*StakingWalletCSTEthDepositEventIterator, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "EthDepositEvent", roundNumRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTEthDepositEventIterator{contract: _StakingWalletCST.contract, event: "EthDepositEvent", logs: logs, sub: sub}, nil
}

// WatchEthDepositEvent is a free log subscription operation binding the contract event 0x725e2879ec3f698d2b2186dc65d2b3c8f4768f1fd80c31be29fe310a83500c98.
//
// Solidity: event EthDepositEvent(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNFTs)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchEthDepositEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTEthDepositEvent, roundNum []*big.Int) (event.Subscription, error) {

	var roundNumRule []interface{}
	for _, roundNumItem := range roundNum {
		roundNumRule = append(roundNumRule, roundNumItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "EthDepositEvent", roundNumRule)
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

// ParseEthDepositEvent is a log parse operation binding the contract event 0x725e2879ec3f698d2b2186dc65d2b3c8f4768f1fd80c31be29fe310a83500c98.
//
// Solidity: event EthDepositEvent(uint256 indexed roundNum, uint256 actionCounter, uint256 depositIndex, uint256 depositId, uint256 depositAmount, uint256 numStakedNFTs)
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseEthDepositEvent(log types.Log) (*StakingWalletCSTEthDepositEvent, error) {
	event := new(StakingWalletCSTEthDepositEvent)
	if err := _StakingWalletCST.contract.UnpackLog(event, "EthDepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletCSTFundTransferFailedIterator is returned from FilterFundTransferFailed and is used to iterate over the raw logs and unpacked data for FundTransferFailed events raised by the StakingWalletCST contract.
type StakingWalletCSTFundTransferFailedIterator struct {
	Event *StakingWalletCSTFundTransferFailed // Event containing the contract specifics and raw log

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
func (it *StakingWalletCSTFundTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletCSTFundTransferFailed)
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
		it.Event = new(StakingWalletCSTFundTransferFailed)
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
func (it *StakingWalletCSTFundTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletCSTFundTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletCSTFundTransferFailed represents a FundTransferFailed event raised by the StakingWalletCST contract.
type StakingWalletCSTFundTransferFailed struct {
	ErrStr      string
	Amount      *big.Int
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundTransferFailed is a free log retrieval operation binding the contract event 0xc4283eec4c85bbabb45475fc62cda5e1b3a382be7ef6b1d14e685b065a9adf74.
//
// Solidity: event FundTransferFailed(string errStr, uint256 amount, address indexed destination)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterFundTransferFailed(opts *bind.FilterOpts, destination []common.Address) (*StakingWalletCSTFundTransferFailedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "FundTransferFailed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTFundTransferFailedIterator{contract: _StakingWalletCST.contract, event: "FundTransferFailed", logs: logs, sub: sub}, nil
}

// WatchFundTransferFailed is a free log subscription operation binding the contract event 0xc4283eec4c85bbabb45475fc62cda5e1b3a382be7ef6b1d14e685b065a9adf74.
//
// Solidity: event FundTransferFailed(string errStr, uint256 amount, address indexed destination)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchFundTransferFailed(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTFundTransferFailed, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "FundTransferFailed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletCSTFundTransferFailed)
				if err := _StakingWalletCST.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
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

// ParseFundTransferFailed is a log parse operation binding the contract event 0xc4283eec4c85bbabb45475fc62cda5e1b3a382be7ef6b1d14e685b065a9adf74.
//
// Solidity: event FundTransferFailed(string errStr, uint256 amount, address indexed destination)
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseFundTransferFailed(log types.Log) (*StakingWalletCSTFundTransferFailed, error) {
	event := new(StakingWalletCSTFundTransferFailed)
	if err := _StakingWalletCST.contract.UnpackLog(event, "FundTransferFailed", log); err != nil {
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
	StakeActionId *big.Int
	TokenId       *big.Int
	StakerAddress common.Address
	NumStakedNFTs *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeActionEvent is a free log retrieval operation binding the contract event 0xcd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef8.
//
// Solidity: event StakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterStakeActionEvent(opts *bind.FilterOpts, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (*StakingWalletCSTStakeActionEventIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "StakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTStakeActionEventIterator{contract: _StakingWalletCST.contract, event: "StakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchStakeActionEvent is a free log subscription operation binding the contract event 0xcd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef8.
//
// Solidity: event StakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchStakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTStakeActionEvent, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "StakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
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

// ParseStakeActionEvent is a log parse operation binding the contract event 0xcd7bda73d32f8fada3eeee8d21563aa19eb2fe86d90b4449cf5252e6f3da7ef8.
//
// Solidity: event StakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs)
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
	StakeActionId *big.Int
	TokenId       *big.Int
	StakerAddress common.Address
	NumStakedNFTs *big.Int
	RewardAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeActionEvent is a free log retrieval operation binding the contract event 0x678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c96.
//
// Solidity: event UnstakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs, uint256 rewardAmount)
func (_StakingWalletCST *StakingWalletCSTFilterer) FilterUnstakeActionEvent(opts *bind.FilterOpts, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (*StakingWalletCSTUnstakeActionEventIterator, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCST.contract.FilterLogs(opts, "UnstakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletCSTUnstakeActionEventIterator{contract: _StakingWalletCST.contract, event: "UnstakeActionEvent", logs: logs, sub: sub}, nil
}

// WatchUnstakeActionEvent is a free log subscription operation binding the contract event 0x678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c96.
//
// Solidity: event UnstakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs, uint256 rewardAmount)
func (_StakingWalletCST *StakingWalletCSTFilterer) WatchUnstakeActionEvent(opts *bind.WatchOpts, sink chan<- *StakingWalletCSTUnstakeActionEvent, stakeActionId []*big.Int, tokenId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

	var stakeActionIdRule []interface{}
	for _, stakeActionIdItem := range stakeActionId {
		stakeActionIdRule = append(stakeActionIdRule, stakeActionIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var stakerAddressRule []interface{}
	for _, stakerAddressItem := range stakerAddress {
		stakerAddressRule = append(stakerAddressRule, stakerAddressItem)
	}

	logs, sub, err := _StakingWalletCST.contract.WatchLogs(opts, "UnstakeActionEvent", stakeActionIdRule, tokenIdRule, stakerAddressRule)
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

// ParseUnstakeActionEvent is a log parse operation binding the contract event 0x678afbb7bbf1c4f3df509d59b51d6e75969703762eb36ed41414dc7c49569c96.
//
// Solidity: event UnstakeActionEvent(uint256 indexed stakeActionId, uint256 indexed tokenId, address indexed stakerAddress, uint256 numStakedNFTs, uint256 rewardAmount)
func (_StakingWalletCST *StakingWalletCSTFilterer) ParseUnstakeActionEvent(log types.Log) (*StakingWalletCSTUnstakeActionEvent, error) {
	event := new(StakingWalletCSTUnstakeActionEvent)
	if err := _StakingWalletCST.contract.UnpackLog(event, "UnstakeActionEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
