// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cosmicgame

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

// IStakingWalletRandomWalkNftMetaData contains all meta data concerning the IStakingWalletRandomWalkNft contract.
var IStakingWalletRandomWalkNftMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numStakerAddresses_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomNumberSeed_\",\"type\":\"uint256\"}],\"name\":\"pickRandomStakerAddressesIfPossible\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStakingWalletRandomWalkNftABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingWalletRandomWalkNftMetaData.ABI instead.
var IStakingWalletRandomWalkNftABI = IStakingWalletRandomWalkNftMetaData.ABI

// IStakingWalletRandomWalkNft is an auto generated Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNft struct {
	IStakingWalletRandomWalkNftCaller     // Read-only binding to the contract
	IStakingWalletRandomWalkNftTransactor // Write-only binding to the contract
	IStakingWalletRandomWalkNftFilterer   // Log filterer for contract events
}

// IStakingWalletRandomWalkNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingWalletRandomWalkNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingWalletRandomWalkNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingWalletRandomWalkNftSession struct {
	Contract     *IStakingWalletRandomWalkNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IStakingWalletRandomWalkNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingWalletRandomWalkNftCallerSession struct {
	Contract *IStakingWalletRandomWalkNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IStakingWalletRandomWalkNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingWalletRandomWalkNftTransactorSession struct {
	Contract     *IStakingWalletRandomWalkNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IStakingWalletRandomWalkNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftRaw struct {
	Contract *IStakingWalletRandomWalkNft // Generic contract binding to access the raw methods on
}

// IStakingWalletRandomWalkNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftCallerRaw struct {
	Contract *IStakingWalletRandomWalkNftCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingWalletRandomWalkNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingWalletRandomWalkNftTransactorRaw struct {
	Contract *IStakingWalletRandomWalkNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingWalletRandomWalkNft creates a new instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNft(address common.Address, backend bind.ContractBackend) (*IStakingWalletRandomWalkNft, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNft{IStakingWalletRandomWalkNftCaller: IStakingWalletRandomWalkNftCaller{contract: contract}, IStakingWalletRandomWalkNftTransactor: IStakingWalletRandomWalkNftTransactor{contract: contract}, IStakingWalletRandomWalkNftFilterer: IStakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// NewIStakingWalletRandomWalkNftCaller creates a new read-only instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftCaller(address common.Address, caller bind.ContractCaller) (*IStakingWalletRandomWalkNftCaller, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftCaller{contract: contract}, nil
}

// NewIStakingWalletRandomWalkNftTransactor creates a new write-only instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingWalletRandomWalkNftTransactor, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftTransactor{contract: contract}, nil
}

// NewIStakingWalletRandomWalkNftFilterer creates a new log filterer instance of IStakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewIStakingWalletRandomWalkNftFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingWalletRandomWalkNftFilterer, error) {
	contract, err := bindIStakingWalletRandomWalkNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftFilterer{contract: contract}, nil
}

// bindIStakingWalletRandomWalkNft binds a generic wrapper to an already deployed contract.
func bindIStakingWalletRandomWalkNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.IStakingWalletRandomWalkNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingWalletRandomWalkNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.contract.Transact(opts, method, params...)
}

// PickRandomStakerAddressesIfPossible is a free data retrieval call binding the contract method 0xe36aee78.
//
// Solidity: function pickRandomStakerAddressesIfPossible(uint256 numStakerAddresses_, uint256 randomNumberSeed_) view returns(address[])
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCaller) PickRandomStakerAddressesIfPossible(opts *bind.CallOpts, numStakerAddresses_ *big.Int, randomNumberSeed_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _IStakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerAddressesIfPossible", numStakerAddresses_, randomNumberSeed_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PickRandomStakerAddressesIfPossible is a free data retrieval call binding the contract method 0xe36aee78.
//
// Solidity: function pickRandomStakerAddressesIfPossible(uint256 numStakerAddresses_, uint256 randomNumberSeed_) view returns(address[])
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) PickRandomStakerAddressesIfPossible(numStakerAddresses_ *big.Int, randomNumberSeed_ *big.Int) ([]common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerAddressesIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, numStakerAddresses_, randomNumberSeed_)
}

// PickRandomStakerAddressesIfPossible is a free data retrieval call binding the contract method 0xe36aee78.
//
// Solidity: function pickRandomStakerAddressesIfPossible(uint256 numStakerAddresses_, uint256 randomNumberSeed_) view returns(address[])
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftCallerSession) PickRandomStakerAddressesIfPossible(numStakerAddresses_ *big.Int, randomNumberSeed_ *big.Int) ([]common.Address, error) {
	return _IStakingWalletRandomWalkNft.Contract.PickRandomStakerAddressesIfPossible(&_IStakingWalletRandomWalkNft.CallOpts, numStakerAddresses_, randomNumberSeed_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Stake(&_IStakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Stake(&_IStakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.StakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.StakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Unstake(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.Unstake(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.UnstakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _IStakingWalletRandomWalkNft.Contract.UnstakeMany(&_IStakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// IStakingWalletRandomWalkNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftStakedIterator struct {
	Event *IStakingWalletRandomWalkNftNftStaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRandomWalkNftNftStaked)
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
		it.Event = new(IStakingWalletRandomWalkNftNftStaked)
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
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRandomWalkNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRandomWalkNftNftStaked represents a NftStaked event raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftStaked struct {
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0x62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletRandomWalkNftNftStakedIterator, error) {

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

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftNftStakedIterator{contract: _IStakingWalletRandomWalkNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0x62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletRandomWalkNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRandomWalkNftNftStaked)
				if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0x62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) ParseNftStaked(log types.Log) (*IStakingWalletRandomWalkNftNftStaked, error) {
	event := new(IStakingWalletRandomWalkNftNftStaked)
	if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingWalletRandomWalkNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftUnstakedIterator struct {
	Event *IStakingWalletRandomWalkNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingWalletRandomWalkNftNftUnstaked)
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
		it.Event = new(IStakingWalletRandomWalkNftNftUnstaked)
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
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingWalletRandomWalkNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingWalletRandomWalkNftNftUnstaked represents a NftUnstaked event raised by the IStakingWalletRandomWalkNft contract.
type IStakingWalletRandomWalkNftNftUnstaked struct {
	ActionCounter *big.Int
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0x08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*IStakingWalletRandomWalkNftNftUnstakedIterator, error) {

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

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingWalletRandomWalkNftNftUnstakedIterator{contract: _IStakingWalletRandomWalkNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0x08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *IStakingWalletRandomWalkNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IStakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingWalletRandomWalkNftNftUnstaked)
				if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0x08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_IStakingWalletRandomWalkNft *IStakingWalletRandomWalkNftFilterer) ParseNftUnstaked(log types.Log) (*IStakingWalletRandomWalkNftNftUnstaked, error) {
	event := new(IStakingWalletRandomWalkNftNftUnstaked)
	if err := _IStakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRandomWalkNftMetaData contains all meta data concerning the StakingWalletRandomWalkNft contract.
var StakingWalletRandomWalkNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"randomWalkNft_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"NftHasAlreadyBeenStaked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callerAddress\",\"type\":\"address\"}],\"name\":\"NftStakeActionAccessDenied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"}],\"name\":\"NftStakeActionInvalidId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"errStr\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actionCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakeActionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numStakedNfts\",\"type\":\"uint256\"}],\"name\":\"NftUnstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"actionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numStakedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numStakerAddresses_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomNumberSeed_\",\"type\":\"uint256\"}],\"name\":\"pickRandomStakerAddressesIfPossible\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomWalkNft\",\"outputs\":[{\"internalType\":\"contractRandomWalkNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId_\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActionIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeActions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwnerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeActionId_\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakeActionIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461004657610019610014610117565b610257565b61002161004b565b6115ee6102f1823960805181818161038c01528181610e6a015261111301526115ee90f35b610051565b60405190565b5f80fd5b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b9061007d90610055565b810190811060018060401b0382111761009557604052565b61005f565b906100ad6100a661004b565b9283610073565b565b5f80fd5b60018060a01b031690565b6100c7906100b3565b90565b6100d3906100be565b90565b6100df816100ca565b036100e657565b5f80fd5b905051906100f7826100d6565b565b906020828203126101125761010f915f016100ea565b90565b6100af565b6101356118df8038038061012a8161009a565b9283398101906100f9565b90565b90565b61014f61014a610154926100b3565b610138565b6100b3565b90565b6101609061013b565b90565b61016c90610157565b90565b90565b61018661018161018b9261016f565b610138565b6100b3565b90565b61019790610172565b90565b60209181520190565b5f7f5468652070726f76696465642061646472657373206973207a65726f2e000000910152565b6101d7601d60209261019a565b6101e0816101a3565b0190565b6101f99060208101905f8183039101526101ca565b90565b61020581610163565b61021f6102196102145f61018e565b6100be565b916100be565b1461022f5761022d90610252565b565b61023761004b565b63eac0d38960e01b81528061024e600482016101e4565b0390fd5b608052565b6102689061026361026a565b6101fc565b565b6102726102d1565b565b5f1b90565b906102855f1991610274565b9181191691161790565b90565b6102a66102a16102ab9261016f565b610138565b61028f565b90565b90565b906102c66102c16102cd92610292565b6102ae565b8254610279565b9055565b6102db5f5f6102b1565b6102ee5f680100000000000000016102b1565b56fe60806040526004361015610013575b61079d565b61001d5f356100cc565b806304fbb3dd146100c75780630d50c189146100c25780632e17de78146100bd57806360294405146100b8578063755b4ef7146100b3578063a2b136fb146100ae578063a694fc3a146100a9578063ca7c1f92146100a4578063e36aee781461009f578063fdbd98b01461009a5763fe939afc0361000e57610769565b610734565b6106e7565b6105e6565b6105a5565b61056d565b610412565b610346565b6102df565b6102ab565b6101f0565b60e01c90565b60405190565b5f80fd5b5f80fd5b5f80fd5b90565b6100f0816100e4565b036100f757565b5f80fd5b90503590610108826100e7565b565b9060208282031261012357610120915f016100fb565b90565b6100dc565b634e487b7160e01b5f52603260045260245ffd5b506801000000000000000090565b90565b6101568161013c565b8210156101705761016860019161014a565b910201905f90565b610128565b1c90565b90565b61018c9060086101919302610175565b610179565b90565b9061019f915461017c565b90565b60016101ad8161013c565b8210156101ca576101c7916101c19161014d565b90610194565b90565b5f80fd5b6101d7906100e4565b9052565b91906101ee905f602085019401906101ce565b565b346102205761021c61020b61020636600461010a565b6101a2565b6102136100d2565b918291826101db565b0390f35b6100d8565b5f80fd5b5f80fd5b5f80fd5b909182601f8301121561026b5781359167ffffffffffffffff831161026657602001926020830284011161026157565b61022d565b610229565b610225565b906020828203126102a1575f82013567ffffffffffffffff811161029c576102989201610231565b9091565b6100e0565b6100dc565b5f0190565b346102da576102c46102be366004610270565b90610816565b6102cc6100d2565b806102d6816102a6565b0390f35b6100d8565b3461030d576102f76102f236600461010a565b610cc2565b6102ff6100d2565b80610309816102a6565b0390f35b6100d8565b680400000000000000026103258161013c565b8210156103425761033f916103399161014d565b90610194565b90565b5f80fd5b346103765761037261036161035c36600461010a565b610312565b6103696100d2565b918291826101db565b0390f35b6100d8565b5f91031261038557565b6100dc565b7f000000000000000000000000000000000000000000000000000000000000000090565b60018060a01b031690565b90565b6103d06103cb6103d5926103ae565b6103b9565b6103ae565b90565b6103e1906103bc565b90565b6103ed906103d8565b90565b6103f9906103e4565b9052565b9190610410905f602085019401906103f0565b565b346104425761042236600461037b565b61043e61042d61038a565b6104356100d2565b918291826103fd565b0390f35b6100d8565b506801000000000000000090565b90565b61046181610447565b82101561047b57610473600391610455565b910201905f90565b610128565b5f1c90565b61049161049691610480565b610179565b90565b6104a39054610485565b90565b60018060a01b031690565b6104bd6104c291610480565b6104a6565b90565b6104cf90546104b1565b90565b68010000000000000002906104e682610447565b81101561051e576104f691610458565b506105025f8201610499565b9161051b6002610514600185016104c5565b9301610499565b90565b5f80fd5b61052b906103ae565b90565b61053790610522565b9052565b60409061056461056b949695939661055a60608401985f8501906101ce565b602083019061052e565b01906101ce565b565b346105a05761059c61058861058336600461010a565b6104d2565b6105939391936100d2565b9384938461053b565b0390f35b6100d8565b346105d3576105bd6105b836600461010a565b610ff5565b6105c56100d2565b806105cf816102a6565b0390f35b6100d8565b6105e35f5f90610194565b90565b34610616576105f636600461037b565b6106126106016105d8565b6106096100d2565b918291826101db565b0390f35b6100d8565b91906040838203126106435780610637610640925f86016100fb565b936020016100fb565b90565b6100dc565b5190565b60209181520190565b60200190565b61066490610522565b9052565b906106758160209361065b565b0190565b60200190565b9061069c61069661068f84610648565b809361064c565b92610655565b905f5b8181106106ac5750505090565b9091926106c56106bf6001928651610668565b94610679565b910191909161069f565b6106e49160208201915f81840391015261067f565b90565b34610718576107146107036106fd36600461061b565b9061128c565b61070b6100d2565b918291826106cf565b0390f35b6100d8565b610731680100000000000000015f90610194565b90565b346107645761074436600461037b565b61076061074f61071d565b6107576100d2565b918291826101db565b0390f35b6100d8565b346107985761078261077c366004610270565b90611380565b61078a6100d2565b80610794816102a6565b0390f35b6100d8565b5f80fd5b5090565b90565b6107bc6107b76107c1926107a5565b6103b9565b6100e4565b90565b634e487b7160e01b5f52601160045260245ffd5b6107e1906100e4565b5f81146107ef576001900390565b6107c4565b9190811015610804576020020190565b610128565b35610813816100e7565b90565b90916108238284906107a1565b5b806108376108315f6107a8565b916100e4565b111561086957610846906107d8565b9061086361085e610859858786916107f4565b610809565b610cc2565b90610824565b50915050565b90565b9061087c906100e4565b9052565b9061088a90610522565b9052565b601f801991011690565b634e487b7160e01b5f52604160045260245ffd5b906108b69061088e565b810190811067ffffffffffffffff8211176108d057604052565b610898565b906108e86108e16100d2565b92836108ac565b565b6108f460606108d5565b90565b9061094661093d60026109086108ea565b9461091f6109175f8301610499565b5f8801610872565b61093761092e600183016104c5565b60208801610880565b01610499565b60408401610872565b565b610951906108f7565b90565b61095e9051610522565b90565b61097561097061097a926107a5565b6103b9565b6103ae565b90565b61098690610961565b90565b60209181520190565b60207f6e7374616b652069742e00000000000000000000000000000000000000000000917f4f6e6c79204e4654206f776e6572206973207065726d697474656420746f20755f8201520152565b6109ec602a604092610989565b6109f581610992565b0190565b916040610a2a929493610a23610a18606083018381035f8501526109df565b9660208301906101ce565b019061052e565b565b5f7f496e76616c6964204e4654207374616b6520616374696f6e2049442e00000000910152565b610a60601c602092610989565b610a6981610a2c565b0190565b9190610a90906020610a88604086018681035f880152610a53565b9401906101ce565b565b1b90565b91906008610ab1910291610aab5f1984610a92565b92610a92565b9181191691161790565b610acf610aca610ad4926100e4565b6103b9565b6100e4565b90565b90565b9190610af0610aeb610af893610abb565b610ad7565b908354610a96565b9055565b5f90565b610b1291610b0c610afc565b91610ada565b565b91906008610b34910291610b2e60018060a01b0384610a92565b92610a92565b9181191691161790565b610b47906103d8565b90565b90565b9190610b63610b5e610b6b93610b3e565b610b4a565b908354610b14565b9055565b5f90565b610b8591610b7f610b6f565b91610b4d565b565b90565b610b9e610b99610ba392610b87565b6103b9565b6100e4565b90565b610bb5610bbb919392936100e4565b926100e4565b8203918211610bc657565b6107c4565b5f1b90565b90610bdc5f1991610bcb565b9181191691161790565b90610bfb610bf6610c0292610abb565b610ad7565b8254610bd0565b9055565b610c1090516100e4565b90565b610c22610c28919392936100e4565b926100e4565b8201809211610c3357565b6107c4565b916020610c59929493610c5260408201965f8301906101ce565b01906101ce565b565b610c64906103d8565b90565b5f80fd5b60e01b90565b5f910312610c7b57565b6100dc565b604090610ca9610cb09496959396610c9f60608401985f85019061052e565b602083019061052e565b01906101ce565b565b610cba6100d2565b3d5f823e3d90fd5b610cdf610cd9680100000000000000028390610458565b5061086f565b90610ce982610948565b9133610d08610d02610cfd60208701610954565b610522565b91610522565b03610f2357610d195f808301610b00565b610d265f60018301610b73565b610dca610d45610d355f610499565b610d3f6001610b8a565b90610ba6565b91610d50835f610be6565b610da55f6002610d73610d6d68040000000000000002889061014d565b90610194565b93610d9f610d8360408b01610c06565b83610d98680100000000000000028990610458565b5001610be6565b01610b00565b610dc468040000000000000002610dbe60408801610c06565b9061014d565b90610ada565b610def610ddf68010000000000000001610499565b610de96001610b8a565b90610c13565b610e028168010000000000000001610be6565b91610e0e5f8501610c06565b903392610e4d610e47610e417f08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc94610abb565b94610abb565b94610b3e565b94610e62610e596100d2565b92839283610c38565b0390a4610e8e7f00000000000000000000000000000000000000000000000000000000000000006103e4565b6323b872dd90610e9d30610c5b565b90610eaa5f339501610c06565b92813b15610f1e575f610ed091610edb8296610ec46100d2565b98899788968795610c6b565b855260048501610c80565b03925af18015610f1957610eed575b50565b610f0c905f3d8111610f12575b610f0481836108ac565b810190610c71565b5f610eea565b503d610efa565b610cb2565b610c67565b50610f3060208301610954565b610f4a610f44610f3f5f61097d565b610522565b91610522565b145f14610f7757610f7390610f5d6100d2565b91829163023df6b160e21b835260048301610a6d565b0390fd5b3390610f9a610f846100d2565b9283926348aca7ef60e11b8452600484016109f9565b0390fd5b90610faf60018060a01b0391610bcb565b9181191691161790565b90610fce610fc9610fd592610b3e565b610b4a565b8254610f9e565b9055565b610fe2906100e4565b5f198114610ff05760010190565b6107c4565b610ffe816114b9565b61102361101368010000000000000001610499565b61101d6001610b8a565b90610c13565b6110368168010000000000000001610be6565b6110aa611056611050680100000000000000028490610458565b5061086f565b611062845f8301610be6565b61106f3360018301610fb9565b6110a561107b5f610499565b9161108b83916002839101610be6565b61109f85916804000000000000000261014d565b90610ada565b610fd9565b906110b5825f610be6565b908233919261110b6110f96110f36110ed7f62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f94610abb565b94610abb565b94610b3e565b946111026100d2565b918291826101db565b0390a46111377f00000000000000000000000000000000000000000000000000000000000000006103e4565b6323b872dd90339061114830610c5b565b9392813b156111bd575f61116f9161117a82966111636100d2565b98899788968795610c6b565b855260048501610c80565b03925af180156111b85761118c575b50565b6111ab905f3d81116111b1575b6111a381836108ac565b810190610c71565b5f611189565b503d611199565b610cb2565b610c67565b606090565b67ffffffffffffffff81116111df5760208091020190565b610898565b906111f66111f1836111c7565b6108d5565b918252565b369037565b9061122561120d836111e4565b9260208061121b86936111c7565b92019103906111fb565b565b600161123391016100e4565b90565b634e487b7160e01b5f52601260045260245ffd5b61125661125c916100e4565b916100e4565b908115611267570690565b611236565b9061127682610648565b811015611287576020809102010190565b610128565b6112946111c2565b5061129d6111c2565b916112a75f610499565b90816112bb6112b55f6107a8565b916100e4565b116112c7575b50505090565b93909192506112d583611200565b92935b846112eb6112e55f6107a8565b916100e4565b1115611374576112fa90611227565b9061136e61135b611355600161134f61134961133a61133461132561131e8b61150c565b8c9061124a565b6804000000000000000261014d565b90610194565b68010000000000000002610458565b5061086f565b016104c5565b966107d8565b95611369869188909261126c565b610880565b906112d8565b509250505f80806112c1565b909161138d8284906107a1565b5b806113a161139b5f6107a8565b916100e4565b11156113d3576113b0906107d8565b906113cd6113c86113c3858786916107f4565b610809565b610ff5565b9061138e565b50915050565b60407f746f206265207374616b6564206f6e6c79206f6e63652e000000000000000000917f54686973204e46542068617320616c7265616479206265656e207374616b65645f8201527f20696e2074686520706173742e20416e204e465420697320616c6c6f7765642060208201520152565b6114596057606092610989565b611462816113d9565b0190565b9190611489906020611481604086018681035f88015261144c565b9401906101ce565b565b156114935750565b6114b59061149f6100d2565b9182916315fdfd0160e31b835260048301611466565b0390fd5b61150a906114ef6114d56114cf6001849061014d565b90610194565b6114e76114e15f6107a8565b916100e4565b14829061148b565b6115046114fc6001610b8a565b91600161014d565b90610ada565b565b61151e90611518610afc565b50611567565b90565b90565b611530611535916100e4565b611521565b9052565b61154581602093611524565b0190565b60200190565b5190565b61155f61156491610480565b610abb565b90565b61158e61159d6115b592611579610afc565b506115826100d2565b92839160208301611539565b602082018103825203826108ac565b6115af6115a98261154f565b91611549565b20611553565b9056fea26469706673582212201bd2f88ba6e878599c3433cb4915df720921575812909618e0b08ef5d81fba7664736f6c634300081c0033",
}

// StakingWalletRandomWalkNftABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingWalletRandomWalkNftMetaData.ABI instead.
var StakingWalletRandomWalkNftABI = StakingWalletRandomWalkNftMetaData.ABI

// StakingWalletRandomWalkNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingWalletRandomWalkNftMetaData.Bin instead.
var StakingWalletRandomWalkNftBin = StakingWalletRandomWalkNftMetaData.Bin

// DeployStakingWalletRandomWalkNft deploys a new Ethereum contract, binding an instance of StakingWalletRandomWalkNft to it.
func DeployStakingWalletRandomWalkNft(auth *bind.TransactOpts, backend bind.ContractBackend, randomWalkNft_ common.Address) (common.Address, *types.Transaction, *StakingWalletRandomWalkNft, error) {
	parsed, err := StakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingWalletRandomWalkNftBin), backend, randomWalkNft_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingWalletRandomWalkNft{StakingWalletRandomWalkNftCaller: StakingWalletRandomWalkNftCaller{contract: contract}, StakingWalletRandomWalkNftTransactor: StakingWalletRandomWalkNftTransactor{contract: contract}, StakingWalletRandomWalkNftFilterer: StakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// StakingWalletRandomWalkNft is an auto generated Go binding around an Ethereum contract.
type StakingWalletRandomWalkNft struct {
	StakingWalletRandomWalkNftCaller     // Read-only binding to the contract
	StakingWalletRandomWalkNftTransactor // Write-only binding to the contract
	StakingWalletRandomWalkNftFilterer   // Log filterer for contract events
}

// StakingWalletRandomWalkNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingWalletRandomWalkNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingWalletRandomWalkNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingWalletRandomWalkNftSession struct {
	Contract     *StakingWalletRandomWalkNft // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StakingWalletRandomWalkNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingWalletRandomWalkNftCallerSession struct {
	Contract *StakingWalletRandomWalkNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// StakingWalletRandomWalkNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingWalletRandomWalkNftTransactorSession struct {
	Contract     *StakingWalletRandomWalkNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// StakingWalletRandomWalkNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftRaw struct {
	Contract *StakingWalletRandomWalkNft // Generic contract binding to access the raw methods on
}

// StakingWalletRandomWalkNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftCallerRaw struct {
	Contract *StakingWalletRandomWalkNftCaller // Generic read-only contract binding to access the raw methods on
}

// StakingWalletRandomWalkNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingWalletRandomWalkNftTransactorRaw struct {
	Contract *StakingWalletRandomWalkNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingWalletRandomWalkNft creates a new instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNft(address common.Address, backend bind.ContractBackend) (*StakingWalletRandomWalkNft, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNft{StakingWalletRandomWalkNftCaller: StakingWalletRandomWalkNftCaller{contract: contract}, StakingWalletRandomWalkNftTransactor: StakingWalletRandomWalkNftTransactor{contract: contract}, StakingWalletRandomWalkNftFilterer: StakingWalletRandomWalkNftFilterer{contract: contract}}, nil
}

// NewStakingWalletRandomWalkNftCaller creates a new read-only instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftCaller(address common.Address, caller bind.ContractCaller) (*StakingWalletRandomWalkNftCaller, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftCaller{contract: contract}, nil
}

// NewStakingWalletRandomWalkNftTransactor creates a new write-only instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingWalletRandomWalkNftTransactor, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftTransactor{contract: contract}, nil
}

// NewStakingWalletRandomWalkNftFilterer creates a new log filterer instance of StakingWalletRandomWalkNft, bound to a specific deployed contract.
func NewStakingWalletRandomWalkNftFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingWalletRandomWalkNftFilterer, error) {
	contract, err := bindStakingWalletRandomWalkNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftFilterer{contract: contract}, nil
}

// bindStakingWalletRandomWalkNft binds a generic wrapper to an already deployed contract.
func bindStakingWalletRandomWalkNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingWalletRandomWalkNftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakingWalletRandomWalkNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingWalletRandomWalkNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.contract.Transact(opts, method, params...)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) ActionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "actionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.ActionCounter(&_StakingWalletRandomWalkNft.CallOpts)
}

// ActionCounter is a free data retrieval call binding the contract method 0xfdbd98b0.
//
// Solidity: function actionCounter() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) ActionCounter() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.ActionCounter(&_StakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) NumStakedNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "numStakedNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.NumStakedNfts(&_StakingWalletRandomWalkNft.CallOpts)
}

// NumStakedNfts is a free data retrieval call binding the contract method 0xca7c1f92.
//
// Solidity: function numStakedNfts() view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) NumStakedNfts() (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.NumStakedNfts(&_StakingWalletRandomWalkNft.CallOpts)
}

// PickRandomStakerAddressesIfPossible is a free data retrieval call binding the contract method 0xe36aee78.
//
// Solidity: function pickRandomStakerAddressesIfPossible(uint256 numStakerAddresses_, uint256 randomNumberSeed_) view returns(address[])
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) PickRandomStakerAddressesIfPossible(opts *bind.CallOpts, numStakerAddresses_ *big.Int, randomNumberSeed_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "pickRandomStakerAddressesIfPossible", numStakerAddresses_, randomNumberSeed_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PickRandomStakerAddressesIfPossible is a free data retrieval call binding the contract method 0xe36aee78.
//
// Solidity: function pickRandomStakerAddressesIfPossible(uint256 numStakerAddresses_, uint256 randomNumberSeed_) view returns(address[])
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) PickRandomStakerAddressesIfPossible(numStakerAddresses_ *big.Int, randomNumberSeed_ *big.Int) ([]common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerAddressesIfPossible(&_StakingWalletRandomWalkNft.CallOpts, numStakerAddresses_, randomNumberSeed_)
}

// PickRandomStakerAddressesIfPossible is a free data retrieval call binding the contract method 0xe36aee78.
//
// Solidity: function pickRandomStakerAddressesIfPossible(uint256 numStakerAddresses_, uint256 randomNumberSeed_) view returns(address[])
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) PickRandomStakerAddressesIfPossible(numStakerAddresses_ *big.Int, randomNumberSeed_ *big.Int) ([]common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.PickRandomStakerAddressesIfPossible(&_StakingWalletRandomWalkNft.CallOpts, numStakerAddresses_, randomNumberSeed_)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) RandomWalkNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "randomWalkNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) RandomWalkNft() (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.RandomWalkNft(&_StakingWalletRandomWalkNft.CallOpts)
}

// RandomWalkNft is a free data retrieval call binding the contract method 0x755b4ef7.
//
// Solidity: function randomWalkNft() view returns(address)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) RandomWalkNft() (common.Address, error) {
	return _StakingWalletRandomWalkNft.Contract.RandomWalkNft(&_StakingWalletRandomWalkNft.CallOpts)
}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) StakeActionIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "stakeActionIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActionIds(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActionIds is a free data retrieval call binding the contract method 0x60294405.
//
// Solidity: function stakeActionIds(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) StakeActionIds(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActionIds(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 nftId, address nftOwnerAddress, uint256 index)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) StakeActions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftId           *big.Int
	NftOwnerAddress common.Address
	Index           *big.Int
}, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "stakeActions", arg0)

	outstruct := new(struct {
		NftId           *big.Int
		NftOwnerAddress common.Address
		Index           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftOwnerAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 nftId, address nftOwnerAddress, uint256 index)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeActions(arg0 *big.Int) (struct {
	NftId           *big.Int
	NftOwnerAddress common.Address
	Index           *big.Int
}, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActions(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// StakeActions is a free data retrieval call binding the contract method 0xa2b136fb.
//
// Solidity: function stakeActions(uint256 ) view returns(uint256 nftId, address nftOwnerAddress, uint256 index)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) StakeActions(arg0 *big.Int) (struct {
	NftId           *big.Int
	NftOwnerAddress common.Address
	Index           *big.Int
}, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeActions(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// UsedNfts is a free data retrieval call binding the contract method 0x04fbb3dd.
//
// Solidity: function usedNfts(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCaller) UsedNfts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingWalletRandomWalkNft.contract.Call(opts, &out, "usedNfts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsedNfts is a free data retrieval call binding the contract method 0x04fbb3dd.
//
// Solidity: function usedNfts(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) UsedNfts(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.UsedNfts(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// UsedNfts is a free data retrieval call binding the contract method 0x04fbb3dd.
//
// Solidity: function usedNfts(uint256 ) view returns(uint256)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftCallerSession) UsedNfts(arg0 *big.Int) (*big.Int, error) {
	return _StakingWalletRandomWalkNft.Contract.UsedNfts(&_StakingWalletRandomWalkNft.CallOpts, arg0)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) Stake(opts *bind.TransactOpts, nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "stake", nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Stake(&_StakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 nftId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) Stake(nftId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Stake(&_StakingWalletRandomWalkNft.TransactOpts, nftId_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) StakeMany(opts *bind.TransactOpts, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "stakeMany", nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeMany(&_StakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// StakeMany is a paid mutator transaction binding the contract method 0xfe939afc.
//
// Solidity: function stakeMany(uint256[] nftIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) StakeMany(nftIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.StakeMany(&_StakingWalletRandomWalkNft.TransactOpts, nftIds_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) Unstake(opts *bind.TransactOpts, stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "unstake", stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Unstake(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 stakeActionId_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) Unstake(stakeActionId_ *big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.Unstake(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionId_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactor) UnstakeMany(opts *bind.TransactOpts, stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.contract.Transact(opts, "unstakeMany", stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.UnstakeMany(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// UnstakeMany is a paid mutator transaction binding the contract method 0x0d50c189.
//
// Solidity: function unstakeMany(uint256[] stakeActionIds_) returns()
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftTransactorSession) UnstakeMany(stakeActionIds_ []*big.Int) (*types.Transaction, error) {
	return _StakingWalletRandomWalkNft.Contract.UnstakeMany(&_StakingWalletRandomWalkNft.TransactOpts, stakeActionIds_)
}

// StakingWalletRandomWalkNftNftStakedIterator is returned from FilterNftStaked and is used to iterate over the raw logs and unpacked data for NftStaked events raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftStakedIterator struct {
	Event *StakingWalletRandomWalkNftNftStaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletRandomWalkNftNftStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRandomWalkNftNftStaked)
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
		it.Event = new(StakingWalletRandomWalkNftNftStaked)
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
func (it *StakingWalletRandomWalkNftNftStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRandomWalkNftNftStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRandomWalkNftNftStaked represents a NftStaked event raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftStaked struct {
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftStaked is a free log retrieval operation binding the contract event 0x62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) FilterNftStaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletRandomWalkNftNftStakedIterator, error) {

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

	logs, sub, err := _StakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftNftStakedIterator{contract: _StakingWalletRandomWalkNft.contract, event: "NftStaked", logs: logs, sub: sub}, nil
}

// WatchNftStaked is a free log subscription operation binding the contract event 0x62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) WatchNftStaked(opts *bind.WatchOpts, sink chan<- *StakingWalletRandomWalkNftNftStaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftStaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRandomWalkNftNftStaked)
				if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
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

// ParseNftStaked is a log parse operation binding the contract event 0x62773741191803b5cec48480156933243e422a1fb1ea9967dab3ee30df2da95f.
//
// Solidity: event NftStaked(uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) ParseNftStaked(log types.Log) (*StakingWalletRandomWalkNftNftStaked, error) {
	event := new(StakingWalletRandomWalkNftNftStaked)
	if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWalletRandomWalkNftNftUnstakedIterator is returned from FilterNftUnstaked and is used to iterate over the raw logs and unpacked data for NftUnstaked events raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftUnstakedIterator struct {
	Event *StakingWalletRandomWalkNftNftUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWalletRandomWalkNftNftUnstaked)
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
		it.Event = new(StakingWalletRandomWalkNftNftUnstaked)
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
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWalletRandomWalkNftNftUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWalletRandomWalkNftNftUnstaked represents a NftUnstaked event raised by the StakingWalletRandomWalkNft contract.
type StakingWalletRandomWalkNftNftUnstaked struct {
	ActionCounter *big.Int
	StakeActionId *big.Int
	NftId         *big.Int
	StakerAddress common.Address
	NumStakedNfts *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNftUnstaked is a free log retrieval operation binding the contract event 0x08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) FilterNftUnstaked(opts *bind.FilterOpts, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (*StakingWalletRandomWalkNftNftUnstakedIterator, error) {

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

	logs, sub, err := _StakingWalletRandomWalkNft.contract.FilterLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakingWalletRandomWalkNftNftUnstakedIterator{contract: _StakingWalletRandomWalkNft.contract, event: "NftUnstaked", logs: logs, sub: sub}, nil
}

// WatchNftUnstaked is a free log subscription operation binding the contract event 0x08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) WatchNftUnstaked(opts *bind.WatchOpts, sink chan<- *StakingWalletRandomWalkNftNftUnstaked, stakeActionId []*big.Int, nftId []*big.Int, stakerAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingWalletRandomWalkNft.contract.WatchLogs(opts, "NftUnstaked", stakeActionIdRule, nftIdRule, stakerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWalletRandomWalkNftNftUnstaked)
				if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
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

// ParseNftUnstaked is a log parse operation binding the contract event 0x08e7047cd3ef25f4a72589ed1d73eb0af1ca8a5957f9539e126dff45870979cc.
//
// Solidity: event NftUnstaked(uint256 actionCounter, uint256 indexed stakeActionId, uint256 indexed nftId, address indexed stakerAddress, uint256 numStakedNfts)
func (_StakingWalletRandomWalkNft *StakingWalletRandomWalkNftFilterer) ParseNftUnstaked(log types.Log) (*StakingWalletRandomWalkNftNftUnstaked, error) {
	event := new(StakingWalletRandomWalkNftNftUnstaked)
	if err := _StakingWalletRandomWalkNft.contract.UnpackLog(event, "NftUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
